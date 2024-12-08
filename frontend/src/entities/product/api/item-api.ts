import { item } from "../model/product-model"

export async function getItems(): Promise<item[]> {
    // В будущем это будет на затычка, а полноценный зпрос
    // const req = await axios.get()

    // Затычка
    const data: item[] = [
        {
            id: 1,
            name: "Кожанка 72 ITAM x Hack Club",
            price: 5200,
        },
        {
            id: 2,
            name: "Брюки ITAM x Hack Club",
            price: 3400,
        },
        {
            id: 3,
            name: "Кожанка 72 ITAM x Hack Club",
            price: 5200,
        },
        {
            id: 4,
            name: "Брюки ITAM x Hack Club",
            price: 3400,
        },
    ]

    return data
}
