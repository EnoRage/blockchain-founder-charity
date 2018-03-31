using BlockChainMobileBack.Models;
using Microsoft.Extensions.Options;
using MongoDB.Bson.Serialization;
using MongoDB.Driver;

namespace BlockChainMobileBack.Repository
{
    public class DbContext
    {
        private readonly IMongoDatabase _database;

        public DbContext(IOptions<Settings> settings)
        {
            var client = new MongoClient(settings.Value.ConnectionString);
            if (client != null)
                _database = client.GetDatabase(settings.Value.Database);
        }

        public IMongoCollection<TransactionHistory> TransactionCollection => _database.GetCollection<TransactionHistory>("transactionHistory");
        public IMongoCollection<User> UsersCollection => _database.GetCollection<User>("users");
        public IMongoCollection<FoundationOptions> FoundationsCollection => _database.GetCollection<FoundationOptions>("foundations");
    }
}