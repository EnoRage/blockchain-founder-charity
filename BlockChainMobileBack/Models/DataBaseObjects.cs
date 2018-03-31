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

        public string Name { get; set; }

        public DateTime FoundedDate { get; set; }

        public double Capital { get; set; }

        public string Country { get; set; }

        public string Mission { get; set; }
    }
    /*
name
foundedDate
capitalisation
country
mission
      */
}
