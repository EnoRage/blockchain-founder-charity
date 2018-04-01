using System;
using System.Collections.Generic;
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
            
        public string EthPrvKey { get; set; } = "";

        public string EthAddress { get; set; } = "";

        public List<Tuple<string, string, string, string>> Foundations { get; set; } 
            = new List<Tuple<string, string, string, string>>();
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
