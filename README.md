# GitHub Inbox Bot

A bot to manage your GitHub notifications directly from Slack.

## Features

-   [x] Receive GitHub notifications in Slack
-   [ ] Mark notifications as read
-   [ ] Snooze notifications for later

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/github-inbox-bot.git
    ```
2. Install dependencies:
    ```sh
    cd github-inbox-bot
    ./build.sh
    ```

## Configuration

1. Create a `.env` file in the root directory and add your GitHub and Slack tokens:
    ```env
    GITHUB_TOKEN=your_github_token
    SLACK_TOKEN=your_slack_token
    SLACK_CHANNEL=your_slack_channel
    ```

## Usage

1. Start the bot:
    ```sh
    ./github-inbox-bot
    ```
2. Invite the bot to your Slack channel and start receiving notifications.
3. <b>IMPORTANT!!!</b> - Make sure the TOKEN has not expired. In case the GitHub Token has expired, manually change the token in al environments. Adject "days-until-expiry" in sendExpiryTokenNotification.go

## Contributing

This Repository is not open for contribution.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any questions or suggestions, please open an issue or contact the repository owner.
