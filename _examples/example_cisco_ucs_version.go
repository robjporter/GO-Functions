package main

import (
	"../cisco/ucs"
	"fmt"
)

func main() {
	ucs.GetWebData()
	fmt.Println("SUGGESTED RELEASES ********************************")
	fmt.Println("SUGGESTED RELEASE:                       >", ucs.GetSuggestedRelease())
	fmt.Println("SUGGESTED 2 RELEASE UPGRADE:             >", ucs.GetSuggestedReleaseTrain("2"))
	fmt.Println("SUGGESTED 2.0 RELEASE UPGRADE:           >", ucs.GetSuggestedReleaseTrain("2.0"))
	fmt.Println("SUGGESTED 2.0(1m) RELEASE UPGRADE:       >", ucs.GetSuggestedReleaseTrain("2.0(1m)"))
	fmt.Println("SUGGESTED 2.2(3e) RELEASE UPGRADE:       >", ucs.GetSuggestedReleaseTrain("2.2(3e)"))
	fmt.Println("SUGGESTED 2.2(8f) RELEASE UPGRADE:       >", ucs.GetSuggestedReleaseTrain("2.2(8f)"))
	fmt.Println("SUGGESTED 3.1(1l) RELEASE UPGRADE:       >", ucs.GetSuggestedReleaseTrain("3.1(1l)"))
	fmt.Println("SUGGESTED 3.2(1l) RELEASE UPGRADE:       >", ucs.GetSuggestedReleaseTrain("3.2(1l)"))

	fmt.Println("LATEST RELEASES    ********************************")
	fmt.Println("LATEST RELEASE:                          >", ucs.GetLatestRelease())
	fmt.Println("LATEST 3.0 RELEASE TRAIN:                >", ucs.GetLatestReleaseTrain("3.0"))
	fmt.Println("LATEST 3.1 RELEASE TRAIN:                >", ucs.GetLatestReleaseTrain("3.1"))
	fmt.Println("LATEST 4.0 RELEASE TRAIN:                >", ucs.GetLatestReleaseTrain("4.0"))
	fmt.Println("LATEST 3.1(1l) RELEASE TRAIN:            >", ucs.GetLatestReleaseTrain("3.1(1l)"))
	fmt.Println("LATEST 3.1(1a) RELEASE TRAIN:            >", ucs.GetLatestReleaseTrain("3.1(1a)"))
	fmt.Println("LATEST 3.1(2f) RELEASE TRAIN:            >", ucs.GetLatestReleaseTrain("3.1(2f)"))
	fmt.Println("DEFERRED RELEASES  ********************************")
	fmt.Println("IS 2.0(1m) DEFERRED RELEASE:             >", ucs.GetIsDeferredRelease("2.0(1m)"))
	fmt.Println("IS 3.1(2f) DEFERRED RELEASE:             >", ucs.GetIsDeferredRelease("3.1(2f)"))
	fmt.Println("SHOW SUGGESTED RELEASES  ***************************")
	ucs.ShowSuggestedReleases()
	fmt.Println("SHOW LATEST RELEASES  ******************************")
	ucs.ShowLatestReleases()
	fmt.Println("SHOW ALL RELEASES  *********************************")
	ucs.ShowAllReleases()
	fmt.Println("SHOW DEFERRED RELEASES  ****************************")
	ucs.ShowDeferredReleases()
	fmt.Println("GET SUGGESTED RELEASES  ****************************")
	fmt.Println(ucs.GetSuggestedReleases())
	fmt.Println("GET LATEST RELEASES  *******************************")
	fmt.Println(ucs.GetLatestReleases())
	fmt.Println("GET ALL RELEASES  **********************************")
	fmt.Println(ucs.GetAllofReleases())
	fmt.Println("GET DEFERRED RELEASES  *****************************")
	fmt.Println(ucs.GetDeferredReleases())

	fmt.Println("FUNCTIONS ******************************************")
	fmt.Println("IS SUGGESTED VERSION:                    >", ucs.IsSuggestedReleaseTrain("3.1(2e)"))
	fmt.Println("IS LATEST VERSION:                       >", ucs.IsLatestReleaseTrain("3.1(2e)"))
	fmt.Println("IS SUGGESTED OR LATEST VERSION:          >", ucs.GetIsSuggestedOrLatest("3.1(2f)"))

}
