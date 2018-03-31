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
        public IMongoCollection<FoundationOptions> FoundationsCollection => _database.GetCollection<FoundationOptions>("Foundations");
        /*public IMongoCollection<Doctor> DoctorCollection => _database.GetCollection<Doctor>("DoctorCollection");
        public IMongoCollection<Pacient> PacientCollection => _database.GetCollection<Pacient>("PacientCollection");
        public IMongoCollection<Dependency> DependenciesCollection => _database.GetCollection<Dependency>("DependencyCollection");
        //public IMongoCollection<Doctor> DoctorCollection => _database.GetCollection<Doctor>("DoctorCollection");*/
    }
}