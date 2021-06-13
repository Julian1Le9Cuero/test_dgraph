
import axios from 'axios'

// Current time in unix format
const ts = Math.floor(Date.now()/1000 - 5000);
const url = 'http://localhost:3000';

class HomeService{
    // Get buyers, transactions and products
    static getAllBuyers(){
        return new Promise(async (resolve, reject) => {
            try {
                const res = await axios.get(`${url}/${ts}`);
                const data = res.data;
                resolve(data.buyers);
            } catch (error) {
                reject(error);
            }
        })
    }

    static getAllProducts(){
        return new Promise(async (resolve, reject) => {
            try {
                const res = await axios.get(`${url}/${ts}`);
                const data = res.data;
                resolve(data.products);
            } catch (error) {
                reject(error);
            }
        })
    }

    static getAllTransactions(){
        return new Promise(async (resolve, reject) => {
            try {
                const res = await axios.get(`${url}/${ts}`);
                const data = res.data;
                resolve(data.transactions);
            } catch (error) {
                reject(error);
            }
        })
    }

    // Get buyers
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