# HOW TO RUN

Create a `config.json` file. Then create an auth config object, this will be imported into the reun and unmarshalled.

## Sample Config object

You can verbatim copy paste this object into your `config.json` but remember to fill out proper credentials depending on your run environment.
Note that the port doesn't have to be 5102 but since this is the root sub service it's recomended to start at 5102 and count up.

```json
{
    "CustomerCfg":{
        "StorageConfig":{
            "ElasticConfig":{
                "ElasticAddr":"",
                "EIndex":"customer",
                "EType":"test-customer_data"
            }
        }
    },
    "Port":5102
}

```
