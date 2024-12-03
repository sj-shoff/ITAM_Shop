import { RootState } from "@app/storage"
import { useSelector } from "react-redux"

export const useAppSelector = useSelector.withTypes<RootState>()
