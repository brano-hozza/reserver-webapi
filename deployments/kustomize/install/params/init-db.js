const mongoHost = process.env.RESERVER_API_MONGODB_HOST
const mongoPort = process.env.RESERVER_API_MONGODB_PORT

const mongoUser = process.env.RESERVER_API_MONGODB_USERNAME
const mongoPassword = process.env.RESERVER_API_MONGODB_PASSWORD

const database = process.env.RESERVER_API_MONGODB_DATABASE

const retrySeconds = parseInt(process.env.RETRY_CONNECTION_SECONDS || "5") || 5;

// try to connect to mongoDB until it is not available
let connection;
while(true) {
    try {
        connection = Mongo(`mongodb://${mongoUser}:${mongoPassword}@${mongoHost}:${mongoPort}`);
        break;
    } catch (exception) {
        print(`Cannot connect to mongoDB: ${exception}`);
        print(`Will retry after ${retrySeconds} seconds`)
        sleep(retrySeconds * 1000);
    }
}

const collections = [
    "doctor", "reservation", "department", "room"
]

const entities = {
    "doctor": [
        {
            "id": "doctor-id-1",
            "name": "Dr. Hrasok Malicky",
            "department": "department-id-1",
        },
        {
            "id": "doctor-id-2",
            "name": "Dr. Janko Hrasko",
            "department": "department-id-1",
        },
        {
            "id": "doctor-id-3",
            "name": "Dr. Jozko Mrkvicka",
            "department": "department-id-2",
        }
    ],
    "department": [
        {
            "id": "department-id-1",
            "name": "Dermatológia"
        },
        {
            "id": "department-id-2",
            "name": "Očná ambulancia"
        }
    ],
    "room": [
        {
            "id": "room-id-1",
            "roomNumber": "Poschodie 1, miestnosť 1",
        },
        {
            "id": "room-id-2",
            "roomNumber": "Poschodie 1, miestnosť 2",
        },
        {
            "id": "room-id-3",
            "roomNumber": "Poschodie 2, miestnosť 1",
        }
    ],
    "reservation": [
        {
            "id": "reservation-id-1",
            "doctor": "doctor-id-1",
            "room": "room-id-1",
            "department": "department-id-1",
        }
    ]
}

const databases = connection.getDBNames()

for(const collection of collections) {
    // if database and collection exists, exit with success - already initialized
    if (databases.includes(database)) {
        const dbInstance = connection.getDB(database)
        const existingCollections = dbInstance.getCollectionNames()
        if (existingCollections.includes(collection)) {
            print(`Collection '${collection}' already exists in database '${database}'`)
            continue;
        }
    }

    // initialize
    // create database and collection
    const db = connection.getDB(database)
    db.createCollection(collection)

    // create indexes
    db[collection].createIndex({ "id": 1 })

    //insert sample data
    let result = db[collection].insertMany(entities[collection]);

    if (result.writeError) {
        console.error(result)
        print(`Error when writing the data: ${result.errmsg}`)
    }
}
// exit with success
process.exit(0);