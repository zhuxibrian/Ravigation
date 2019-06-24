package cmd

import (
	"github.com/spf13/cobra"
	"Ravigation/ravictl/handler"
	"Ravigation/ravigation/storage"
)

var deviceType string
var deviceName string
var deviceHost string
var devicePort uint32

var x, y, z float64
var pointIndex int

var node, btn, cmdName string

var cmdGroup string



func init() {
	addCmd := newAddCmd()
	rootCmd.AddCommand(addCmd)
}


func newAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [device]",
		Short: "add device",
		Long:  `add device with device info`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "connectInfo":
				params := storage.ConnectInfo{
					Device: deviceType,
					Name: deviceName,
					Host: deviceHost,
					Port: devicePort,
				}
				handler.PostConnectInfo(params)
			case "point":
				point := storage.Point{
					X:x,
					Y:y,
					Z:z,
				}
				handler.PostPoint(point, pointIndex)
			case "btncmdmap":
				btncmdmap := storage.ButtonCmdGroup{
					NodeName:node,
					ButtonName:btn,
					CmdGroupName:cmdName,
				}
				handler.PostButtonCmdGroup(btncmdmap)
			case "cmd":
				handler.PostCmdGroup(cmdGroup)

			}

		},
	}

	cmd.Flags().StringVarP(&deviceType, "type", "t", "", "device type")
	cmd.Flags().StringVarP(&deviceName, "name", "n", "", "device name")
	cmd.Flags().StringVarP(&deviceHost, "host", "", "", "device host")
	cmd.Flags().Uint32VarP(&devicePort, "port", "p", 0, "device port")

	cmd.Flags().IntVarP(&pointIndex, "index", "i", 0, "point index")
	cmd.Flags().Float64VarP(&x, "x", "x", 0, "point x coordinate")
	cmd.Flags().Float64VarP(&y, "y", "y", 0, "point y coordinate")
	cmd.Flags().Float64VarP(&z, "z", "z", 0, "point z coordinate")

	cmd.Flags().StringVarP(&node, "nodename", "o", "", "node name")
	cmd.Flags().StringVarP(&btn, "btnname", "b", "", "button name")
	cmd.Flags().StringVarP(&cmdName, "cmdname", "m", "", "command name")

	cmd.Flags().StringVarP(&cmdGroup, "cmd", "c", "", "command group")
	return cmd
}



