using System;
using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace BlockChainMobileBack.Models
{

    public class TestInt
    {
        [BsonRepresentation(BsonType.ObjectId)]
        [BsonIgnoreIfDefault]
        [BsonIgnoreIfNull]
        [BsonId]
        public string Id { get; set; }

        public int Value { get; set; } = 0;
    }

    public class FoundationOptions
    {
        [BsonRepresentation(BsonType.ObjectId)]
        [BsonIgnoreIfDefault]
        [BsonIgnoreIfNull]
        [BsonId]
        public string Id { get; set; }

        public string Name { get; set; } = "";
        //gggg
        public int FoundedDate { get; set; } = 0;

        public float Capital { get; set; } = 0.0F;

        public string Country { get; set; } = "";

        public string Mission { get; set; } = "";
    }

    public class User
    {
        [BsonRepresentation(BsonType.ObjectId)]
        [BsonIgnoreIfDefault]
        [BsonIgnoreIfNull]
        [BsonId]
        public string Id { get; set; }

        public string Name { get; set; } = "";

        public string Password { get; set; } = "";

        public string PubKey { get; set; } = "";

        public string PrvKey { get; set; } = "";

        public string Adress { get; set; } = "";
    }

    public class TransactionHistory
    {
        [BsonRepresentation(BsonType.ObjectId)]
        [BsonIgnoreIfDefault]
        [BsonIgnoreIfNull]
        [BsonId]
        public string Id { get; set; }

        public string UserId { get; set; }

        public string OrgId { get; set; }

        //gggg.mm.dd*hh:mm:ss
        public string DateTime { get; set; }

        public float Summ { get; set; }
    }
    /*
name
foundedDate
capitalisation
country
mission
      */
}
