import { Button, Input, Card } from "@nextui-org/react";
import { useState } from "react";
import { Product } from "@pages/admin/index";

interface ProductFormProps {
  onAddProduct: (product: Product) => void;
}

const ProductForm: React.FC<ProductFormProps> = ({ onAddProduct }) => {
  const [name, setName] = useState<string>("");
  const [price, setPrice] = useState<number | "">("");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (name && typeof price === "number") {
      onAddProduct({ name, price });
      setName("");
      setPrice("");
    }
  };

  return (
    <Card>
      <form onSubmit={handleSubmit}>
        <Input
          label="Название товара"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <Input
          label="Цена"
          type="number"
          // value={price}
          onChange={(e) => setPrice(Number(e.target.value))}
        />
        <Button type="submit">Добавить</Button>
      </form>
    </Card>
  );
};

export default ProductForm;
