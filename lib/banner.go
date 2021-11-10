package lib

import (
	"fmt"

	"github.com/binganao/breachkit/pkg/logger"
)

func banner() {
	title := "" +
		"\n" +
		" ██████╗ ██████╗ ███████╗ █████╗  ██████╗██╗  ██╗██╗  ██╗██╗████████╗\n" +
		" ██╔══██╗██╔══██╗██╔════╝██╔══██╗██╔════╝██║  ██║██║ ██╔╝██║╚══██╔══╝\n" +
		" ██████╔╝██████╔╝█████╗  ███████║██║     ███████║█████╔╝ ██║   ██║   \n" +
		" ██╔══██╗██╔══██╗██╔══╝  ██╔══██║██║     ██╔══██║██╔═██╗ ██║   ██║   \n" +
		" ██████╔╝██║  ██║███████╗██║  ██║╚██████╗██║  ██║██║  ██╗██║   ██║   \n" +
		" ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝   ╚═╝   \n"
	version := "\tbreachkit version: 0.1 release  author: bingan"
	github := "\tgithub: https://github.com/binganao/breachkit\n"
	fmt.Println(logger.LightWhite(title))
	fmt.Println(logger.White(version))
	fmt.Println(logger.White(github))
}
