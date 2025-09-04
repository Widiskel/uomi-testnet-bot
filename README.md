# Uomi Testnet BOT

## Table Of Contents

- [Uomi Testnet BOT](#uomi-testnet-bot)
  - [Table Of Contents](#table-of-contents)
  - [Prerequisite](#prerequisite)
  - [Uomi Incentive Testnet](#uomi-incentive-testnet)
  - [BOT FEATURE](#bot-feature)
  - [Setup \& Configure BOT](#setup--configure-bot)
    - [Linux](#linux)
  - [Update Bot](#update-bot)
  - [NOTE](#note)
  - [CONTRIBUTE](#contribute)
  - [SUPPORT](#support)

## Prerequisite

- Git
- Go version >= go1.24.5

## Uomi Incentive Testnet

#New

Uomi Testnet + Synthra
Reward : Confirmed
Network : Uomi Testnet

Link: https://synthra.org/0KM7VHLZ

INVITE CODE : `0KM7VHLZ`

- Connect Wallet (New / BURNER)
- Enter Invite Code
- Daily TX
- Go To Portofolio & Connect X
- Done!

## BOT FEATURE

- Multi Account
- Support PK and Seed
- Daily Swap

## Setup & Configure BOT

### Linux

1. clone project repo
   ```
   git clone https://github.com/Widiskel/uomi-testnet-bot.git
   cd uomi-testnet-bot
   ```
2. Prepare template file
   ```
   cp .env.example .env
   cp config/accounts_tmp.json config/accounts.json
   ```
3. Configure .env
   ```
   nano .env
   ```
4. Configure accounts
   ```
   nano config/accounts.json
   ```
5. Install dependency
   ```
   go mod tidy
   ```
6. Run the bot
   ```
   go run cmd/bot/main.go
   ```

## Update Bot

To update bot follow this step :

1. run
   ```
   git pull
   ```
   or
   ```
   git pull --rebase
   ```
   if error run
   ```
   git stash && git pull
   ```
2. run
   ```
   go mod tidy
   ```
3. start the bot

## NOTE

DWYOR & Always use a new wallet when running the bot, I am not responsible for any loss of assets.

## CONTRIBUTE

Feel free to fork and contribute adding more feature thanks.

## SUPPORT

Want to support me for creating another bot ?
**star** my repo or buy me a coffee on

BYBIT UID : `140173364`

BINANCE UID : `39357434`
