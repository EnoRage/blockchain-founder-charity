using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using BlockChainMobileBack.Models;
using MongoDB.Bson;
using MongoDB.Driver;
using Remotion.Linq.Clauses.ResultOperators;

namespace BlockChainMobileBack.Repository
{
    public interface IDataBaseRepository
    {
        //Task<TestInt> CheckAuthorization(string snils, DateTime birthday);
        Task AddFoundationOption(FoundationOptions fo);
        Task<IEnumerable<FoundationOptions>> GetFoundations();
    }
}