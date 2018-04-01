using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using BlockChainMobileBack.Models;
using BlockChainMobileBack.Repository;
using Microsoft.AspNetCore.Mvc;

// For more information on enabling Web API for empty projects, visit https://go.microsoft.com/fwlink/?LinkID=397860

namespace BlockChainMobileBack.Controllers
{
    [Route("api/[controller]")]
    public class UsersController : Controller
    {
        private readonly IDataBaseRepository _rep;
        public UsersController(IDataBaseRepository dataBaseRepository)
        {
            _rep = dataBaseRepository;
        }

        [HttpGet("{log}")]
        public async Task<bool> CheckReg(string log)
        {
           return await _rep.CheckUserReg(log);
        }

        [HttpGet("{log}/{inf}")]
        public async Task<User> CheckReg(string log, string inf)
        {
            return await _rep.GetUser(log);
        }

        [HttpPost("{log}/{pass}")]
        public async Task<string> AddUser(string log)
        {
            await _rep.AddUser(log);
            return "User has been added";
        }
    }
}
