db.auth("admin","admin")
db.createUser(
    {
        user: "sun",
        pwd: "sun#sun",
        roles: [ { role: "readWrite", db: "db_sun" }
        ]
    }
)