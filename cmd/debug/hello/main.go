/*
cobraのコマンドを試しに使ってみる
helloコマンドを追加してみる
helloを実行するとHello, World!と表示される
-h, --helpでヘルプを表示できる
-v, --versionでバージョンを表示できる
Worldの部分を変更できるようにする
そのときのオプションは-w, --whoで指定する
*/
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var who string

// バージョン情報
const version = "1.0.0"

func main() {
	// Cobraのルートコマンドを定義
	var rootCmd = &cobra.Command{
		Use:   "hello",
		Short: "Prints Hello, World!",
		Long:  "This application prints 'Hello, World!' or allows customization of 'World'",
		Run: func(cmd *cobra.Command, args []string) {
			if who == "" {
				who = "World"
			}
			fmt.Printf("Hello, %s!\n", who)
		},
	}

	// --who/-w オプションを追加
	rootCmd.Flags().StringVarP(&who, "who", "w", "", "Specify the target to greet (default: World)")

	// -v/--version オプションを追加してバージョンを表示
	rootCmd.Flags().BoolP("version", "v", false, "Prints the version of the application")
	rootCmd.PreRun = func(cmd *cobra.Command, args []string) {
		if v, _ := cmd.Flags().GetBool("version"); v {
			fmt.Printf("Version: %s\n", version)
			os.Exit(0)
		}
	}

	// エラーハンドリングのためにExecuteを実行
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
