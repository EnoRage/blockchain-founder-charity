using System.Collections.Generic;
using System.Threading.Tasks;
using BlockChainMobileBack.Models;
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

        public async Task<IEnumerable<FoundationOptions>> GetFoundations()
        {
            var filter = Builders<FoundationOptions>.Filter.Empty;

            return await _context.FoundationsCollection.Find(filter).ToListAsync();
        }
    }
}