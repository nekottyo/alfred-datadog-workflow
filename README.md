# alfred-datadog-workflow

[![version](https://img.shields.io/github/v/tag/nekottyo/alfred-datadog-workflow?sort=semver)](https://github.com/nekottyo/alfred-datadog-workflow/releases)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/fe1cb90a9803401cb98c0b95fdcd3f93)](https://app.codacy.com/manual/nekottyo/alfred-datadog-workflow?utm_source=github.com&utm_medium=referral&utm_content=nekottyo/alfred-datadog-workflow&utm_campaign=Badge_Grade_Dashboard)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/nekottyo/alfred-datadog-workflow.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/nekottyo/alfred-datadog-workflow/alerts/)
[![action-bumpr supported](https://img.shields.io/badge/bumpr-supported-ff69b4?logo=github&link=https://github.com/haya14busa/action-bumpr)](https://github.com/haya14busa/action-bumpr)

## Installation
* [Download the latest release](https://github.com/nekottyo/alfred-datadog-workflow/releases)
* Open the downloaded file in Finder

## Usage

To use it, activate Alfred and type in `dd`.

### authentication
register ApiKey and AppKey

#### api key
`dd auth <ENTER>`  
`ApiKey <ENTER>`  
`<INPUT API KEY>`  

#### app key
`dd auth <ENTER>`  
`AppKey <ENTER>`  
`<INPUT APP KEY>`  

### open service

`dd open <services>`

ex) `event`, `monitor`, `dashboard`..

see [config/service.yaml](./config/service.yaml)

### serach monitor

`dd monitor <query>`

### serach dashboard

`dd dashboard <query>`

## License
[MIT](https://github.com/nekottyo/alfred-datadog-workflow/blob/master/LICENSE)

## author
[@nekottyo](https://github.com/nekottyo)
