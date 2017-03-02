# Tweetstorm

Command line tool that takes a text of arbitrary length as input and produces a stream of tweets from it.

## Installation

Download this repository

## Usage

1. A user access token (OAuth1) grants a consumer application access to a user's  Twitter resources.

   Setup an OAuth1 `http.Client` with the consumer key and secret and oauth token and secret. 

    * export TWITTER_CONSUMER_KEY="{your_key}"
    * export TWITTER_CONSUMER_SECRET="{your_key}"
    * export TWITTER_ACCESS_TOKEN="{your_token}"
    * export TWITTER_ACCESS_SECRET="{your_token}"

2. Now, you can use it like this:
   `tweetstorm "Go (often referred to as golang) is a free and open source programming language created at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson. It is a compiled, statically typed language in the tradition of Algol and C, with garbage collection, limited structural typing, memory safety features and CSP-style concurrent programming features added."
` where it will tweet the value of the command line argument.
