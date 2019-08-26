## latest

Enhancement:

* Now we have specific errors for:
  If you're checking those errors before (since they have no error type, you might be checking their message), pay attention that these errors are TencentCloudSDKError now.
    * Network IO error: ClientError.NetworkError, when fail to get response, such as proxy problem, connection timeout. Note that ssl certificate issue is included as well, even though it is not a good choice, we don't have so many cases to handle and neither so much effort to category them.


## 3.0.81 (2019-08-12)

Enhancement:

* Now we have specific errors for:
  If you're checking those errors before (since they have no error type, you might be checking their message), pay attention that these errors are TencentCloudSDKError now.
    * Network IO error: ClientError.IOError, when fail to read response body.
    * Http status code error: ClientError.HttpStatusCodeError, all non 200 status code will return this error.
    * Json parsing error: ClientError.ParseJsonError, when fail to parse json content from response body.

## 3.0.75 (2019-07-26)

fixes:

* correctly return error when fail to flat paramter to url

enhancement:

* use bytes.Buffer instead of string concate for performance issue

## 3.0.10 (2018-07-06)

* update apis
    * aai
    * cdb
    * cvm

## 3.0.9 (2018-07-02)

* update apis
    * aai
    * cbs
    * cws
    * postgres
* add new product
    * dts

## 3.0.8 (2018-06-21)

* update apis
	* batch
	* cbs
	* cvm
	* ds
* add predefined region info

## 3.0.7 (2018-06-14)

* update apis
	* batch
	* cdb
	* cvm
	* cws
	* vpc

## 3.0.6 (2018-06-08)

* api new
	* cis
	* ds
	* ms
* api update
	* aai
	* cbs
	* cvm
	* vpc
* fix float field conversion problem

## 3.0.5 (2018-05-31)

update api to latest:

* cbs
* cdb
* cvm
* vpc
