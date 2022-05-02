import { useZIndex } from "element-plus/lib";
const { nextZIndex } = useZIndex();

/**
 * @description 当前最大 z-index
 */
export const maxZIndex = nextZIndex();
