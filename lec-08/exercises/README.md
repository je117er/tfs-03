# Overview
## MySQL vs. ElasticSearch querying performance comparison
* For the phrase "donald trump could've made america great again", ElasticSearch took **22X** less time than MySQL to find for a match:
![mysql vs. elastic search querying performance](ESvsMySQL/query_benchmark_results.jpg "MySQL vs. ElasticSearch querying performance")
* Rerun on the same phrase ended up with a much faster time for ES probably due to automatic caching. The more frequently a search term is used, the less time ES will take to return a response.
* There's some discrepancy between how go's *Time* measures the elapsed time in comparison with elastic's built-in *TookInMillis*. Might totally have been due to own poor implementation, however.
## CRUD against database *people*
### ER diagram
![people database](people/db_diagram.jpg "ER Diagram for database people")
### CRUD Operations
* GET /people
* GET /person/:id
* POST /person
* PUT /person/:id
* DELETE /person/:id