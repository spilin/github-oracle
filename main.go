package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/big"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	oracle "github-oracle/contracts"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/go-github/github"
	"github.com/spf13/viper"
)

type Config struct {
	Endpoint   string `mapstructure:"endpoint"`
	OracleEOA  string `mapstructure:"oracle_eoa"`
	PrivateKey string `mapstructure:"private_key"`
	Query      string `mapstructure:"query"`
}

var (
	config Config
)

type PRData struct {
	User   string
	Repo   string
	Number int
}

func main() {
	configPath := flag.String("config", "config/local.yaml", "config/local.yaml")
	flag.Parse()
	viper.SetConfigFile(*configPath)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("github-oracle: Could not load environment variables", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("github-oracle: Could not load environment variables", err)
	}

	interrupt := make(chan os.Signal, 10)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGINT)

	client, err := ethclient.Dial(config.Endpoint)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(config.OracleEOA)
	oracleInstance, err := oracle.NewOracle(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{},
	}

	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	blocks := make(map[string]bool)
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			if vLog.Topics[0].Hex() != config.Query {
				continue
			}
			query, err := oracleInstance.OracleFilterer.ParseQuery(vLog)
			log.Println(query.Raw.Topics)
			if err != nil {
				log.Println(err)
				continue
			}
			if blocks[query.Id.String()] {
				log.Println("Dupe")
				continue
			}
			blocks[query.Id.String()] = true
			prUrl, err := parseUrl(query.PrUrl)

			if err != nil {
				log.Println(err)
				continue
			}
			result, err := checkPullRequest(prUrl, query.GhUsername)
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Println(result)
			err = respond(client, oracleInstance, query.Id, result)
			if err != nil {
				log.Println(err)
				continue
			}
		case <-interrupt:
			os.Exit(0)
		}
	}

}

// Queries github api to fetch PR. If Pr merged - checks if github user made any commits there.
func checkPullRequest(prData PRData, commiter string) (bool, error) {
	ghclient := github.NewClient(nil)
	ctx := context.Background()
	pull, _, err := ghclient.PullRequests.Get(ctx, prData.User, prData.Repo, prData.Number)
	if err != nil {
		return false, err
	}

	if *pull.Merged {
		opts := &github.ListOptions{Page: 1}
		commits, _, err := ghclient.PullRequests.ListCommits(ctx, prData.User, prData.Repo, prData.Number, opts)
		if err != nil {
			return false, err
		}
		keys := make(map[string]bool)
		for _, commit := range commits {
			login := *commit.Committer.Login
			if _, val := keys[login]; !val {
				keys[login] = true
			}
		}
		return keys[commiter], nil
	} else {
		return false, nil
	}
}

// Unpacks githup pull request url into username, repo and pr number
func parseUrl(prUrl string) (PRData, error) {
	var prData PRData
	u, err := url.Parse(prUrl)
	if err != nil {
		return prData, err
	}
	if u.Host != "github.com" {
		return prData, errors.New("not a github URL")
	}

	components := strings.Split(u.Path, "/")
	if components[3] != "pull" {
		return prData, errors.New("not a Pull Request URL")
	}
	if len(components) < 5 {
		return prData, errors.New("missing Pull Request Number")
	}
	prData.User = components[1]
	prData.Repo = components[2]
	prData.Number, err = strconv.Atoi(components[4])

	if err != nil {
		return prData, errors.New("error during conversion")
	}
	return prData, nil
}

// Call a contract to emit Response event with status of PR.
func respond(client *ethclient.Client, oracleInstance *oracle.Oracle, id *big.Int, result bool) error {
	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	// chainId, err := client.ChainID(context.Background())
	// if err != nil {
	// return err
	// }

	// fmt.Println(chainId)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1313161555))
	if err != nil {
		return err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice

	tx, err := oracleInstance.Respond(auth, id, result)
	if err != nil {
		return err
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
	return nil
}
