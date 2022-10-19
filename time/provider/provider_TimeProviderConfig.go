package provider


type TimeProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/time#alias TimeProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

