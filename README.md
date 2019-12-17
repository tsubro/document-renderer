# Document Renderer


> This is a REST based API and utility library to convert various doucment to pdf.
> I made this project with the intention to learn golang and about open source contribution.
> Code is still in pre-mature state and I will be obliged if folks here can support me in this learning journey.

Supported Format are :
* PDF
* JPEG
* GIF
* TIF
* PNG
* DOCX
* XLSX
* HTML

[![Build Status](http://img.shields.io/travis/badges/badgerbadgerbadger.svg?style=flat-square)](https://travis-ci.org/badges/badgerbadgerbadger) [![Dependency Status](http://img.shields.io/gemnasium/badges/badgerbadgerbadger.svg?style=flat-square)](https://gemnasium.com/badges/badgerbadgerbadger) [![Coverage Status](http://img.shields.io/coveralls/badges/badgerbadgerbadger.svg?style=flat-square)](https://coveralls.io/r/badges/badgerbadgerbadger) [![Code Climate](http://img.shields.io/codeclimate/github/badges/badgerbadgerbadger.svg?style=flat-square)](https://codeclimate.com/github/badges/badgerbadgerbadger) [![Github Issues](http://githubbadges.herokuapp.com/badges/badgerbadgerbadger/issues.svg?style=flat-square)](https://github.com/badges/badgerbadgerbadger/issues) [![Pending Pull-Requests](http://githubbadges.herokuapp.com/badges/badgerbadgerbadger/pulls.svg?style=flat-square)](https://github.com/badges/badgerbadgerbadger/pulls) [![Gem Version](http://img.shields.io/gem/v/badgerbadgerbadger.svg?style=flat-square)](https://rubygems.org/gems/badgerbadgerbadger) [![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org) [![Badges](http://img.shields.io/:badges-9/9-ff6799.svg?style=flat-square)](https://github.com/badges/badgerbadgerbadger)

---
## To run the code
> go run main.go

## Sample Request (Optional)

```javascript
URL :http://localhost:8080/job-api
POST
{
  "inputs" : [
    {
      	"input_location":"https://4.img-dpreview.com/files/p/E~TS590x0~articles/3925134721/0266554465.jpeg",
 		"input_mime_type" :"image/jpeg"
    },
     {
      	"input_location":"http://www.africau.edu/images/default/sample.pdf",
 		"input_mime_type" :"application/pdf"
    },
     {
      	"input_location":"http://mtskheta.gov.ge/public/img/1530793528.xlsx",
 		"input_mime_type" :"application/xlsx"
    },
    {
      	"input_location":"https://calibre-ebook.com/downloads/demos/demo.docx",
 		"input_mime_type" :"application/docx"
       
    }
  ],
  "output_location":"http://www.google.com",
  "output_header" : [
    {"name": "name-test", "value":"value-test"},
    {"name": "name-test2", "value":"value-test2"}
  ],
  "output_mime_type" :"application/pdf"
}
```

---

### Setup

> update and install this package first
> Install Mongo-DB
> Install Libra Office 

```shell
$ brew update

```
---

## Contributing

> To get started...

### Step 1

- **Option 1**
    - 🍴 Fork this repo!

- **Option 2**
    - 👯 Clone this repo to your local machine 

### Step 2

- **HACK AWAY!** 🔨🔨🔨

### Step 3

- 🔃 Create a new pull request 
---

## License

[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org)

- **[MIT license](http://opensource.org/licenses/mit-license.php)**

