package main

import ( 
	"fmt"
	"flag"
	"os"
)
/*

wpgo create --php="7.2" --db="mariadb" <projectname> <projectname.local> 
wpgo start <projectname>
wpgo stop <projectname>
*/
func main() {
	createCMD := flag.NewFlagSet("create", flag.ExitOnError)
	PHP_VERSION := createCMD.String("php", "7.1", "PHP Version")
	DB := createCMD.String("db", "mariadb", "Choose Database")
	
	startCMD := flag.NewFlagSet("start", flag.ExitOnError)
	stopCMD := flag.NewFlagSet("stop", flag.ExitOnError)

	if(len(os.Args) < 2) {
		fmt.Println("wpgo requires an action to work!");
		return;
	}
	command := os.Args[1];

	switch(command) {
	case "create":
		createCMD.Parse(os.Args[2:]);
		project := createCMD.Arg(0);
		domain  := createCMD.Arg(1);
		fmt.Println("Creating Project", project, domain, *PHP_VERSION, *DB);
	case "start":
		startCMD.Parse(os.Args[2:]);
		project := startCMD.Arg(0);
		fmt.Println("Starting up Project", project);
	case "stop":
		stopCMD.Parse(os.Args[2:]);
		project := stopCMD.Arg(0);
		fmt.Println("Shutting down Project", project);
	default: 
		fmt.Println("Action doesn't exist! Try one of the ones below")
	}
}
