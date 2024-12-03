import { useDispatch } from "react-redux"
import { AppDispatch } from "@app/storage"

export const useAppDispatch = useDispatch.withTypes<AppDispatch>
