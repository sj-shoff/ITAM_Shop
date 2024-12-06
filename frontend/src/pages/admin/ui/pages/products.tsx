import AppSidebar from "../components/sidebar";
import Header from "../components/header";
import ProductForm from "../components/productform";
import { useState } from "react";
import { Product } from "@pages/admin/index";

const Products: React.FC = () => {
  const [products, setProducts] = useState<Product[]>([]);

  const handleAddProduct = (product: Product) => {
    setProducts([...products, product]);
  };

  return (
    <>
      <Header />
      <AppSidebar />
      <main>
        <ProductForm onAddProduct={handleAddProduct} />
        <ul>
          {products.map((product, index) => (
            <li key={index}>
              {product.name} - {product.price} ₽
            </li>
          ))}
        </ul>
      </main>
    </>
  );
};

export default Products;
