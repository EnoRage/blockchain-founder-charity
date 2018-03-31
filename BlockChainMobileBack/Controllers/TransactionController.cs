using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using BlockChainMobileBack.Models;
using BlockChainMobileBack.Repository;
using Microsoft.AspNetCore.Mvc;

namespace BlockChainMobileBack.Controllers
{
    [Route("api/[controller]")]
    public class TransactionController : Controller
    {
        private readonly IDataBaseRepository _dataBaseRepository;

        public TransactionController(IDataBaseRepository dataBaseRepository)
        {
            _dataBaseRepository = dataBaseRepository;
        }

        [HttpGet("{id}")]
        public async Task<IEnumerable<TransactionHistory>> GetUsersTransaction(string id)
        {
            return await _dataBaseRepository.GetUsersTransaction(id);
        }

        // POST api/values
        [HttpPost("{userId}/{orgId}/{sum}")]
        public async Task<string> Post(string userId, string orgId, float sum)
        {
            await _dataBaseRepository.AddUsersTransaction(userId, orgId, sum);
            return "Sucess Transaction";
        }
    }
}
