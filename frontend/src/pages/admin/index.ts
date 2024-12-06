export interface Product {
    name: string;
    price: number;
  }
  
  export interface User {
    id: number;
    name: string;
    role: string;
    email: string;
  }
  
  export interface StatsData {
    labels: string[];
    sales: number[];
    visits: number[];
  }
  