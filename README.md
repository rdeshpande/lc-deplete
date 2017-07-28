# lc-deplete

Automatically remove available funds from your LendingClub account and withdraw into your bank account.

This is a script designed to run periodically (e.g. via cron) to deplete your account.

## Setup

1) Login to lendingclub.com and proceed to Settings to get an API key.

2) Set the LC_KEY and LC_ACCOUNT_ID environment variables accordingly, and invoke the script.

The script will then either withdraw available funds or exit if you have a zero balance.
