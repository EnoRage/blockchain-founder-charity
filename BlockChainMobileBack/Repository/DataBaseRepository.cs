using System.Collections.Generic;
using System.Threading.Tasks;
using BlockChainMobileBack.Models;
using System.Linq;
using Microsoft.Extensions.Options;
using MongoDB.Driver;

namespace BlockChainMobileBack.Repository
{
    public class DataBaseRepository : IDataBaseRepository
    {
        private readonly DbContext _context;

        public DataBaseRepository(IOptions<Settings> settings)
        {
            _context = new DbContext(settings);
        }

        public async Task AddFoundationOption(FoundationOptions fo)
        {
            await _context.FoundationsCollection.InsertOneAsync(fo);
        }

        public async Task AddUser(string id, string pass)
        {
            await _context.UsersCollection.InsertOneAsync(new User()
            {
                Name = id,
                Password = pass
            });
        }

        public async Task<bool> CheckUserReg(string id, string pass)
        {
            bool result = false;
            var andFilter = Builders<User>.Filter.And(
                Builders<User>.Filter.Eq("Name", id),
                Builders<User>.Filter.Eq("Password", pass));

            result = (await _context.UsersCollection.Find(andFilter).ToListAsync()).FirstOrDefault() != null;
            return result;
        }

        public async Task<IEnumerable<FoundationOptions>> GetFoundations()
        {
            var filter = Builders<FoundationOptions>.Filter.Empty;

            List<FoundationOptions> tmp = await _context.FoundationsCollection.Find(filter).ToListAsync();
            return tmp;
        }
    }
}