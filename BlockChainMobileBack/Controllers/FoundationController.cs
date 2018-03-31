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
    public class FoundationController : Controller
    {
        private readonly IDataBaseRepository _db;

        public FoundationController(IDataBaseRepository rep)
        {
            _db = rep;
        }

        [HttpGet]
        public async Task<IEnumerable<FoundationOptions>> Get()
        {
            return await _db.GetFoundations();
        }

        [HttpPost("{Name}/{FoundedDate}/{Capital}/{Country}/{Mission}")]
        public async Task<string> PostAdd(string Name, int FoundedDate, float Capital, string Country, string Mission)
        {
            await _db.AddFoundationOption(new FoundationOptions()
            {
                Name = Name,
                FoundedDate = FoundedDate,
                Country = Country,
                Capital = Capital,
                Mission = Mission
            });
            return "Sucess Added " + Name + " Foundation";
        }

        /*
         * [HttpPost("{login}/{pass}")]
        public async Task<Doctor> Post(string login, string pass)
        {
            return await _dbRepository.GetDoctor(login, pass);
        }
         */

    }
}
