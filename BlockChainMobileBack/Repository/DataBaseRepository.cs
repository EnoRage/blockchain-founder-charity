using System.Collections.Generic;
using System.Threading.Tasks;
using BlockChainMobileBack.Models;
using System.Linq;
using Microsoft.Extensions.Options;
using MongoDB.Driver;
using System;

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
            var andFilter = Builders<User>.Filter.And(
                Builders<User>.Filter.Eq("Name", id),
                Builders<User>.Filter.Eq("Password", pass));

            return (await _context.UsersCollection.Find(andFilter).ToListAsync())
                    .FirstOrDefault() != null;
        }

        public async Task<IEnumerable<FoundationOptions>> GetFoundations()
        {
            var filter = Builders<FoundationOptions>.Filter.Empty;
            return await _context.FoundationsCollection.Find(filter).ToListAsync(); ;
        }

        public async Task<IEnumerable<TransactionHistory>> GetUsersTransaction(string usersId)
        {
            var filter = Builders<TransactionHistory>.Filter.Eq("UserId", usersId);
            return await _context.TransactionCollection.Find(filter).ToListAsync();
        }

        private static string DateTimeConverter(DateTime inp)
        {
            //gggg.mm.dd*hh:mm:ss
            return $"{inp.Year}.{inp.Month}.{inp.Day}*{inp.Hour}:{inp.Minute}:{inp.Second}";
        }

        public async Task AddUsersTransaction(string usersId, string orgId, float sum)
        {
            TransactionHistory tmp = new TransactionHistory()
            {
                UserId = usersId,
                OrgId = orgId,
                DateTime = DateTimeConverter(DateTime.UtcNow),
                Summ = sum
            };
            await _context.TransactionCollection.InsertOneAsync(tmp);
        }
    }
}