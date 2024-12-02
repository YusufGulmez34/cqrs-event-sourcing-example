set -m

    /entrypoint.sh couchbase-server &

    sleep 15

    # Setup initial cluster/ Initialize Node
    couchbase-cli cluster-init -c 127.0.0.1 --cluster-name $CLUSTER_NAME --cluster-username $COUCHBASE_ADMINISTRATOR_USERNAME \
    --cluster-password $COUCHBASE_ADMINISTRATOR_PASSWORD --services data,index,query,fts --cluster-ramsize 512 --cluster-index-ramsize 512 \
    --cluster-fts-ramsize 512 --index-storage-setting default 

    fg 1