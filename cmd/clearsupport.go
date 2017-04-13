package cmd

import (
	"github.com/fredericlemoine/gotree/io"
	"github.com/spf13/cobra"
)

// clearsupportCmd represents the clearsupport command
var clearsupportCmd = &cobra.Command{
	Use:   "supports",
	Short: "Clear supports from input trees",
	Long:  `Clear supports from input trees.`,
	Run: func(cmd *cobra.Command, args []string) {
		f := openWriteFile(outtreefile)
		treefile, treechan := readTrees(intreefile)
		defer treefile.Close()
		for t := range treechan {
			if t.Err != nil {
				io.ExitWithMessage(t.Err)
			}
			t.Tree.ClearSupports()
			f.WriteString(t.Tree.Newick() + "\n")
		}
		f.Close()
	},
}

func init() {
	clearCmd.AddCommand(clearsupportCmd)
}
