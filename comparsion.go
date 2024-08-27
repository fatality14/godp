package godp

// BranchComparisonResult contains branch comparsion results of same architecture packages
type BranchComparisonResult struct {
	Arch               string        `json:"arch"`
	InSecondNotInFirst []PackageInfo `json:"in_second_not_in_first"`
	InFirstNotInSecond []PackageInfo `json:"in_first_not_in_second"`
	HigherInFirst      []PackageInfo `json:"higher_in_first"`
}

// ComparePackages performs package list comparsion for every present architecture
func ComparePackages(first, second APIResponse) []BranchComparisonResult {
	extractArch := func(resp APIResponse) map[string][]PackageInfo {
		packages := make(map[string][]PackageInfo)
		for _, pkg := range resp.Packages {
			packages[pkg.Arch] = append(packages[pkg.Arch], pkg)
		}
		return packages
	}

	firstPkgsAll := extractArch(first)
	secondPkgsAll := extractArch(second)

	comparisonResults := []BranchComparisonResult{}

	// range all packages by architectures
	for arch, firstPkgs := range firstPkgsAll {
		secondPkgs := secondPkgsAll[arch]
		result := BranchComparisonResult{Arch: arch}

		// make maps for fast package search by name
		firstMap := make(map[string]PackageInfo)
		for _, pkg := range firstPkgs {
			firstMap[pkg.Name] = pkg
		}

		secondMap := make(map[string]PackageInfo)
		for _, pkg := range secondPkgs {
			secondMap[pkg.Name] = pkg
		}

		// compare if first branch package present in second branch
		for _, firstPkg := range firstPkgs {
			secondPkg, exists := secondMap[firstPkg.Name]
			if !exists {
				result.InFirstNotInSecond = append(result.InFirstNotInSecond, firstPkg)
			} else {
				// if present compare versions and release
				if firstPkg.Version > secondPkg.Version || (firstPkg.Version == secondPkg.Version && firstPkg.Release > secondPkg.Release) {
					result.HigherInFirst = append(result.HigherInFirst, firstPkg)
				}
			}
		}

		// add second branch packages that not present in first branch
		for _, secondPkg := range secondPkgs {
			if _, exists := firstMap[secondPkg.Name]; !exists {
				result.InSecondNotInFirst = append(result.InSecondNotInFirst, secondPkg)
			}
		}

		comparisonResults = append(comparisonResults, result)
	}

	return comparisonResults
}
