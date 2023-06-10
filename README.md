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
Charts, Insights and Audiences; both faved and not faved. (*)
```
$ ./scripts/curl/getassets.sh pretty

  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  7761    0  7761    0     0  1901k      0 --:--:-- --:--:-- --:--:-- 2526k
{
  "success": true,
  "message": "count: 30",
  "data": [
    {
      "ID": "32f92457-de96-4f0f-bfd1-9a382a198fd2",
      "Name": "Chart 1",
      "Type": "chart",
      "Data": {
        "ID": "32f92457-de96-4f0f-bfd1-9a382a198fd2",
        "Name": "Chart 1",
        "Title": "Title 1",
        "XAxisTitle": "X-Axis Title 1",
        "YAxisTitle": "Y-Axis Title 1",
        "Data": "ezEuMCwyLjAsMy4wfQ==",
        "Favorite": false
      }
    },
    ...
```

### Fav an asset
```
$ sh scripts/curl/favunfav.sh fav 32f92457-de96-4f0f-bfd1-9a382a198fd2
{"success":true,"data":{}}
```

### Unfav an asset
```
$ sh scripts/curl/favunfav.sh unfav 32f92457-de96-4f0f-bfd1-9a382a198fd2
{"success":true,"data":{}}
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

## Notes

Please review the extended README [here](docs/readme.md) for comprehensive information about Asset Keeper requirements.  
