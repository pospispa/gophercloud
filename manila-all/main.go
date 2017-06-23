package main

import (
	"fmt"
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	sfsapi "github.com/gophercloud/gophercloud/openstack/sharedfilesystems/apiversions"
)

func main() {
	regionName := os.Getenv("OS_REGION_NAME")

	authOpts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		fmt.Printf("AuthOptionsFromEnv failed: (%v)", err)
		fmt.Println("")
		return
	}
	fmt.Println("")
	fmt.Printf("AuthOptionsFromEnv: (%v)", authOpts)
	fmt.Println("")
	provider, err := openstack.AuthenticatedClient(authOpts)
	if err != nil {
		fmt.Printf("AuthenticatedClient failed: (%v)", err)
		fmt.Println("")
		return
	}
	fmt.Println("")
	fmt.Printf("Provider client: (%v)", provider)
	fmt.Println("")
	client, err := openstack.NewSharedFileSystemV2(provider, gophercloud.EndpointOpts{Region: regionName})
	if err != nil {
		fmt.Printf("NewSharedFileSystemV2 failed: (%v)", err)
		fmt.Println("")
		return
	}
	fmt.Printf("Service client: (%v)", client)
	fmt.Println("")

	listRes := sfsapi.List(client)
	fmt.Println("")
	fmt.Printf("List res: (%v)", listRes)
	fmt.Println("")
	allPages, err := listRes.AllPages()
	if err != nil {
		fmt.Println("")
		fmt.Printf("all pages error: %q", err.Error())
		fmt.Println("")
	} else {
		fmt.Println("")
		fmt.Printf("List res, all pages : (%v)", allPages)
		fmt.Println("")
	}

}
