# Asset Keeper

Asset Keeper is a platform that provides users with access to a list of assets. Users can create a personal list of favorites, consisting of assets they want to quickly access from their frontpage dashboard. The following types of assets are supported:

1. **Chart**: A chart with a small title, axes titles, and data.
2. **Insight**: A small piece of text that provides an insight into a specific topic, such as "40% of millennials spend more than 3 hours on social media daily."
3. **Audience**: A series of characteristics describing a specific audience. For this exercise, the focus is on gender (Male, Female), birth country, age groups, hours spent daily on social media, and the number of purchases last month. For example, "Males from 24-35 that spend more than 3 hours on social media daily."

## Key Features

This example is not trivial but rather associated with code that could be found in a real production service. However, its implementation aims to avoid over-engineering and instead focuses on meeting specific requirements. The project follows a clear initial organization, although it does not account for every single corner case.

Similar but relatively more complex examples can be found at the following links:

* [DDD Sphere](https://github.com/orgs/dddsphere/repositories) 
* [Foorester](https://github.com/orgs/foorester/repositories)

Examples of libraries commonly used in various projects include [Zerolog](https://github.com/rs/zerolog) for logging, [Viper](https://github.com/spf13/viper) for 12-Factor configured apps, and [Chi](https://github.com/go-chi/chi) as a routing solution. However, in this case, we have deliberately preferred to prioritize the use of the Go Standard Library whenever possible. It is worth mentioning that the standard library offers production-ready functionality and is widely adopted in real-world implementations.

We have strived to minimize reliance on external dependencies by leveraging the features provided by the Go Standard Library that in their usage, share conceptual similarities.

## Usage
### Run the app
```
$ make run
```

### Get all assets
Charts, Insights and Audiences; both faved and not faved. 
Here, `pretty` uses jq to format the output (*) and `1` and `14` are the page and page size respectively.

```
$ ./scripts/curl/getassets.sh pretty 1 14
Request URL: http://localhost:8080/api/v1/c03dc326-7160-4b63-ac36-7105a4c96fa3/assets?page=1&size=14
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  4727    0  4727    0     0   332k      0 --:--:-- --:--:-- --:--:--  329k
{
  "success": true,
  "count": 14,
  "pages": 2,
  "data": [
    {
      "ID": "146012c3-5fac-491f-8764-84741e223231",
      "Type": "insight",
      "UserID": "",
      "AssetID": "146012c3-5fac-491f-8764-84741e223231",
      "Name": "Insight 10",
      "Description": "⭐ Insight 10",
      "Data": {
        "ID": "146012c3-5fac-491f-8764-84741e223231",
        "Name": "Insight 10",
        "Text": "Text 10",
        "Topic": "Topic 10"
      }
    },
    {
      "ID": "15372e01-2ff5-4ec1-a2b8-9cb183933606",
      "Type": "chart",
      "UserID": "",
      "AssetID": "15372e01-2ff5-4ec1-a2b8-9cb183933606",
      "Name": "johndoe-chart-1",
      "Description": "Faved chart by johndoe",
      "Data": {
        "ID": "15372e01-2ff5-4ec1-a2b8-9cb183933606",
        "Type": "chart",
        "Title": "Title 10",
        "XAxisTitle": "X-Axis Title 10",
        "YAxisTitle": "Y-Axis Title 10",
        "Data": "{123, 49, 48, 44, 50, 48, 44, 51, 48, 125}"
      }
    },
    ...
```

### Fav an asset
```
$./scripts/curl/favunfav.sh fav insight 146012c3-5fac-491f-8764-84741e223231
Request URL: http://localhost:8080/api/v1/efd8cec6-3e45-4fb1-b0d7-3a1be9cfae2c/assets/146012c3-5fac-491f-8764-84741e223231
JSON Body: {
  "type": "insight",
  "action": "fav",
  "name": "Asset Name",
  "description": "Asset Description"
}
{"success":true,"count":1,"data":""}
```

### Unfav an asset
```
$./scripts/curl/favunfav.sh unfav insight 146012c3-5fac-491f-8764-84741e223231
Request URL: http://localhost:8080/api/v1/efd8cec6-3e45-4fb1-b0d7-3a1be9cfae2c/assets/146012c3-5fac-491f-8764-84741e223231
JSON Body: {
  "type": "insight",
  "action": "unfav",
  "name": "Asset Name",
  "description": "Asset Description"
}
{"success":true,"count":1,"data":""}

```

### Update name of a faved asset
```
$ scripts/curl/updatefavedname.sh
{"success":true,"data":{}}
```

(*) If jq is not installed in your system remove the `pretty` flag.
You can also opt installing `jq` first.
This is how you do it:

* Debian / Ubuntu
```
$ sudo aptitude install jq
```

* Arch / Manjaro
```
$ sudo pacman -Sy jq
```

* RedHat / CentOS
```
$ yum install epel-release -y
...
$ yum update -y 
...
$ yum install jq -y
...
```

## Considerations
* A database [migrator](https://github.com/adrianpk/migration) could be helpful for database initialization.
At the moment all the script for database initial setup reside here:

* [Migrations](/scripts/sql/pg/migrations)
* [Fixtures](/scripts/sql/pg/fixtures)

## Todo
* Improve test coverage.
* Add Godoc comments.

## Notes

Please review the extended README [here](docs/readme.md) for comprehensive information about Asset Keeper requirements.  
