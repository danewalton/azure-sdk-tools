$secrets = @{}
$secretsDir = "/mnt/secrets/static/*"
Get-ChildItem -Path $secretsDir | ForEach-Object {
    foreach($line in Get-Content $_) {
        $idx = $line.IndexOf("=")
        if ($idx -gt 0) {
            $key = $line.Substring(0, $idx)
            $val = $line.Substring($idx + 1)
            $secrets.Add($key, $val)
        }
    }
}

mkdir /azure
Copy-Item "/scripts/stress-test/test-resources-post.ps1" -Destination "/azure/"
Copy-Item "/mnt/testresources/*" -Destination "/azure/"

# Capture output so we don't print environment variable secrets
$env = & /common/TestResources/New-TestResources.ps1 `
    -BaseName $env:RESOURCE_GROUP_NAME `
    -ResourceGroupName $env:RESOURCE_GROUP_NAME `
    -SubscriptionId $secrets.AZURE_SUBSCRIPTION_ID `
    -TenantId $secrets.AZURE_TENANT_ID `
    -ProvisionerApplicationId $secrets.AZURE_CLIENT_ID `
    -ProvisionerApplicationSecret $secrets.AZURE_CLIENT_SECRET `
    -TestApplicationId $secrets.AZURE_CLIENT_ID `
    -TestApplicationSecret $secrets.AZURE_CLIENT_SECRET `
    -TestApplicationOid $secrets.AZURE_CLIENT_OID `
    -Location 'westus2' `
    -DeleteAfterHours 168 `
    -ServiceDirectory '/azure/' `
    -SuppressVsoCommands:$true `
    -CI `
    -Force

#>
