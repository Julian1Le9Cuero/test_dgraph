import axios from 'axios'

// Current time in unix format
const url = 'http://localhost:3000';
const today = Math.floor(Date.now()/1000)

class HomeService{
    // Get buyers, transactions and products
    static getAllData(filters){
        const {date, limit, offset, model} = filters
        let filterDate = date
        if (filterDate === null) {
            filterDate = today
        }
        return new Promise(async (resolve, reject) => {
            try {
                console.log(`${url}/home?date=${filterDate}&limit=${limit}&offset=${offset}&model=${model}`)
                const res = await axios.get(`${url}/home?date=${filterDate}&limit=${limit}&offset=${offset}&model=${model}`);
                const data = res.data;
                resolve(data);
            } catch (error) {
                reject(error);
            }
        })
    }


    // Get buyers endpoint
    static getBuyers(){
        return new Promise(async (resolve, reject) => {
            try {
                const res = await axios.get(`${url}/buyers`);
                const data = res.data;
                resolve(data);
            } catch (error) {
                reject(error);
            }
        })  
    }

    // Get buyer by id
    static getBuyerById(id){
        return new Promise(async (resolve, reject) => {
            try {
                const res = await axios.get(`${url}/buyers/${id}`);
                const data = res.data;
                resolve(data);
            } catch (error) {
                reject(error);
            }
        })
    }
}

export default HomeService