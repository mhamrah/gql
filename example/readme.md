yab -t thrift/starwars.thrift -s starwars --method Query::Human --peer localhost:28941 --request '{ "request": {"id":"1000"} }'

curl -XPOST -H 'Rpc-Caller: mlh' -H 'Rpc-Procedure: human' -H 'Rpc-Service: starwars' -H 'Context-TTL-MS: 10000' -H 'Rpc-Encoding: json' localhost:24034 -d '{ "id": "1000" }'

curl  -H  'Rpc-Caller: mlh'  -H 'Context-TTL-MS: 10000'  localhost:24035 -d 'query FetchLukeQuery {
                        Human(ID: "1000") {
                                name
                                id
                        }
                }'
