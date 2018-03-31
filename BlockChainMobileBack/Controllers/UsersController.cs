using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
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
        // GET: api/values

        // GET api/values/5
        [HttpGet("{log}/{pass}")]
        public async Task<bool> CheckReg(string log, string pass)
        {
           return await _rep.CheckUserReg(log, pass);
        }

        [HttpPost("{log}/{pass}")]
        public async Task<string> AddUser(string log, string pass)
        {
            await _rep.AddUser(log, pass);
            return "User has been added";
        }
    }
}
