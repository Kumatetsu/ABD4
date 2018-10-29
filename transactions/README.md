--To add the index to elastic search directly from the this directory:------

curl -X PUT localhost:9200/name_of_new_index -d @transaction_index.json

--To verify if the index was correctly inserted in ES and in detail:

curl -X GET "localhost:9200/name_of_new_index/_mapping/transaction

--List all indexes:

curl -X GET "localhost:9200/_cat/indices?v"


