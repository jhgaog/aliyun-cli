# CHANGE LOG

### Master

### 3.0.73

- update: meta data

### 3.0.72

- update: meta data

### 3.0.71

- update: meta data

### 3.0.70

- update: meta data
### 3.0.70

- update: meta data
### 3.0.70

- update: meta data

### 3.0.69

- update: meta data
### 3.0.67

- update: meta data
### 3.0.67

- update: meta data
### 3.0.67

- update: meta data

### 3.0.66

- update: meta data

### 3.0.65

- update: meta data

### 3.0.64

- update: meta data

### 3.0.63

- update: meta data

### 3.0.62

- update: meta data

### 3.0.61

- update: meta data

### 3.0.60

- update: meta data

### 3.0.59

- update: meta data

### 3.0.58

- update: meta data

### 3.0.57

- update: meta data

### 3.0.56

- update: meta data

### 3.0.55

- update: meta data
- fix: the upper limit of array support
- unhide `--skip-secure-verify`

### 3.0.54

- update: meta data

### 3.0.53

- update: meta data

### 3.0.52

- update: meta data

### 3.0.51

- update: meta data

### 3.0.50

- update: meta data

### 3.0.49

- update: meta data

### 3.0.48

- update: meta data

### 3.0.47

- update: meta data
- fix: zsh completion script

### 3.0.46

- update: meta data

### 3.0.45

- update: meta data

### 3.0.44

- update: meta data

### 3.0.43

- update: meta data
- feature: Automatic type recognition `AK`,`StsToken`,`RamRoleArn`,`RsaKeyPair`,`EcsRamRole`

### 3.0.42

- update: meta data
- add: environment variables `ALIBABACLOUD_REGION_ID`, `ALICLOUD_REGION_ID`

### 3.0.41

- update: meta data
- fix: unable to get accessKeyID from environment variable

### 3.0.40

- update: meta data
- update: go-bindata program 
from `https://github.com/jteeuwen/go-bindata` to `https://github.com/shuLhan/go-bindata`
- add: environment variables 
`ALIBABACLOUD_ACCESS_KEY_ID`, `ALICLOUD_ACCESS_KEY_ID`, `ALIBABACLOUD_ACCESS_KEY_SECRET`, `ALICLOUD_ACCESS_KEY_SECRET`
- fix: cancel HTML character escaping

### 3.0.39

- update: meta data
- update: auto-complete case-insensitive

### 3.0.38

- update: meta data
- update: ossutil 1.6.11
- fix: Skip https certificate verification

### 3.0.37

- update: meta data

### 3.0.36

- update: meta data

### 3.0.35

- update: meta data

### 3.0.34

- update: meta data

### 3.0.33

- update: meta data
- update: support oss log

### 3.0.32

- update: meta data

### 3.0.31

- update: meta data
- update: flag `--retry-timeout` renamed to `--read-timeout`
- add: flag `--connect-timeout`
- fix: flag `--retry-timeout` bug

### 3.0.30

- update: meta data
- add: credential type `RamRoleArnWithRoleName`
  >This type is to add RamRoleArn credentials based on EcsRamRole, and does not store information such as ak locally. When the ecs role has the AssumeRole permission, it can be used to exchange permissions for another role.
```shell
$ aliyun configure --mode RamRoleArnWithRoleName --profile ecsarn
Configuring profile 'ecsarn' in 'RamRoleArnWithRoleName' authenticate mode...
Ecs Ram Role []: <YourEcsRamRole>
Ram Role Arn []: <YourRamRoleArn>
Role Session Name []: <YourRoleSessionName>
Expired Seconds [900]: 
Default Region Id []: cn-hangzhou
Default Output Format [json]: json (Only support json)
Default Language [zh|en] en:
Saving profile[ecsarn] ...Done
```
- add: flag `--expired-seconds` to specify expiration time
- update: Help information for the configure command

### 3.0.29

- update: meta data
- update: Cancel the scientific notation of numbers
- fix: Invoke error when reading parameter value from file using `-FILE`

### 3.0.28

- update: meta data
- fix: AccessKeySecret in the `configure set` command is the same as AccessKeyId

### 3.0.27

- update: meta data
- fix: Restful interface reported invalid url error

### 3.0.26

- update: meta data
- fix: help generating information bug
- Upgrade the ossutil component to version 1.6.6
- Package manager switches to Go Modules

### 3.0.25

- update: ros, default version 2019-09-10

### 3.0.24

- update: arms, default version 2019-08-08

### 3.0.23

- fix: fix `--pager` error
- add: product sae

### 3.0.22

- update: API meta data

### 3.0.21

- update: API meta data
- update: dependence

### 3.0.20

- add product: fnf, default version 2019-03-15
- new: support for reading the value of a parameter from a file
    >*Note: Only works on api parameters and rpc style*  
    
    >The value of some parameters is very long, which is not conducive to terminal input.
    >You can add `-FILE` to the original parameter to specify to read from the file, 
    >followed by the file path.

    >example:  
    >command `aliyun ecs CreateInstance` has a parameter `--UserData`, You can specify the value of this parameter from a file by using `--UserData-FILE <filePath>`
    ```shell
    aliyun ecs CreateInstance --UserData-FILE '/home/Document/user_data'
    ```
    >*NOTE: The above example is not a complete command, just a demonstration of how new features are used.*

### 3.0.19

- add product: oos, default version 2019-06-01

### 3.0.18

- updata: API meta data
- fix: the prompt information of the delete command
- improve: command 'configure' can automatically configure the 
  default profile when no profile name is specified 
- fix: cannot force calls to products whose metadata is not entered
- fix: output message mistake 

### 3.0.17

- update meta data
- update error message when configuration fails
- Forced calls to support different API styles
  
### 3.0.16

- Support environment variable `ALIBABA_CLOUD_VENDOR`
  > Set environment variables `ALIBABA_CLOUD_VENDOR` to add a custom UserAgent ID
- Update API meta data, and update endpoint pattern of product cr
- Update gosdk dependence
- Change logic of flag `--pager`

### 3.0.15

- Update API meta data, update version of product CMS to 2019-01-01

### 3.0.14

- Update API meta data and increase kms product help
- Remove the `RootFilter[0]` input at field `rows=` that confuses the user
> The previous `rows` field requires the user to enter `RootFilter[0]`, which will cause confusion for the user and has been deleted.  
>example:
>   >`aliyun ecs DescribeRegions --output cols=RequestId`  

>   >`aliyun ecs DescribeRegions --output cols=RegionId rows=Regions.Region`  
- flag `--pager` can specified collections path
- add row number for flag `--output` output format.
> If you want show row number at output format, you can use field `num=ture` after flag `--output` to enable the num.

### 3.0.13

- update API meta data
- update flag `--output`
> The filtering is based on array form, use `cols=` to Specify the fields to be filtered, and `rows=` to specify the array where the specified field is located, And if the field is below the root element, you can omit `rows=`.
>   >`aliyun ecs DescribeRegions --output cols=RequestId`  

>   >`aliyun ecs DescribeRegions --output cols=RegionId rows=RootFilter[0].Regions.Region`  

>NOTE: Fields in two arrays cannot be mixed

### 3.0.12

- update API meta data
- update dependency package
- added go1.12 test environment

### 3.0.11

- update API meta data
- improve output format of command `configure get`
- add `--config-path` flag to specify file path which has a single profile
>(this flag can use with other flags, like `--access-key-id`, and flag value priority is higher than file content.)  
- add `--retry-timeout` and `--retry-count` flags to support timeout setting 

### v3.0.10

- update open api meta
- support json pretty output

### v3.0.1 Final Version

- support `version` command
- fix `--dryrun` flag bugs


### v3.0.0 GA Version

- refactoring cli package design, support composite flag with fields
- refactoring openapi package design, make it more extensible
- support `--quiet` flag
- support `--dryrun` flag
- support `aliyun oss --profile xxx`

### 0.81

- support `--output`
- support `--waiter`
- use `go -ldflags` to enable single Version in Makefile

### 0.80

- support auto completion for zsh/bash
- fix bugs for RepeatList parameter
- refactoring RamRoleArn and EcsRamRole authenticate flow
- oss command can support RamRoleArn and EcsRamRole authenticate mode
- oss command can support --profile and other configure flags

### 0.70

- integrate `ossutil` toolset with aliyun-cli
- optimize `--help` command messages
- config flags (such as ak, profile, sts) can used with openapi call
- support `configure delete`
- fix bug with restful force call

### 0.61

- support --all-pages flags to merge pager APIs

### 0.60

- support suggestions
- optimized error and help message
- integrate more completion of metadata
- fix some caller bugs

### 0.50

- support i18n `aliyun-openapi-meta`
- full support `configure [get|set|list]` command
- optimize help
- support `--quiet` flag

### 0.33

- fix bug for error processing when rpc/restful call
- auto add Content-Type header for restful call

### 0.32

- auto migrate legacy settings

### 0.31

- fix bug of check parameters, skip Action, Region parameters
- support `aliyun configure list` command

### 0.30

- integrate with 64 products meta
- implemented help command for product and api
- support fully certificated mode AK|StsToken|RamRoleArn|EcsRamRole|RsaKeyPair,

### 0.16

- support --content-type flag to set Header
- support --body-file flat to use file as body input

### 0.15

- support ecs-ram-role
- fix cross platform build problem
- test after configure

### 0.12

- fix bug for configure
- ignore case of ProductName

### 0.11

- Support simple ROA call

### 0.1

- Refactoring with golang
- Basic configure
- Auto endpoint locator
- 2018.1.11