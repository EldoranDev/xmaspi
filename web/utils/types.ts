type RenderElementType = 'animation' | 'static';

interface RenderElement {
    name: string
    displayName?: string
    description?: string
    type: RenderElementType
}