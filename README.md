# GoShipIt
Application to easily post "ShipIt Squirrels" to Slack via slash commands. Written in Go.

# Configuration

Create a slack.json file and place it in the same folder as the GoShipIt executable. The contents of this configuration file can be customized:

```
{
	"bot_username":"GoShipit",
	"webhook_url":"http://yourApiUrl.GetThisFromSlack/somekey"
}
```

Then configure GoShipIt in the Configure Integrations section of Slack.

## Slash Command Integration

In Slack, go to Home > Integrations and add a slash command intigration for GoShipIt with the following settings:

**Command:** /goshipit

**URL:** http://{yourdomain}/slash

**Method:** POST

## Incoming Webhook Integration

Add a new incoming webhook integration with the following settings:

**Channel:** #general

Take the value of the Webhook URL field and place that url in the `webhook_url` section of the slack.json file in GoShipIt's root directory.

# Running GoShipit

Use this command to run GoShipIt:

`$ ./goshipit {yourdomain.com} {yourport} {use https true/false}`

# Usage

* `/goshipit` will randomly select any ShipIt squirrel and display it in Slack.

* `/goshipit categories` will show what the available categories are and how many ShipIt squirrel images are in them.

* `/goshipit {categoryname}` will randomly select a ShipIt squirrel image from the specified category.