
export function useXmaspi() {
    const apiUrl = useState('xmaspi-apiUrl', () => "/api")

    async function getElements(): Promise<Array<RenderElement>> {
        const [{ data: animations }, { data: statics } ] = await Promise.all([
            useFetch<Array<RenderElement>>(`${apiUrl.value}/animations`),
            useFetch<Array<RenderElement>>(`${apiUrl.value}/statics`),
        ]);

        const elements: Array<RenderElement> = [
            ...(animations.value!.map((a) => ({...a, type: 'animation' as RenderElementType}))),
            ...(statics.value!.map((s) => ({...s, type: 'static' as RenderElementType}))),
        ];

        elements.sort((a, b) => a.name.localeCompare(b.name));

        return elements;
    }

    async function apply(element: RenderElement): Promise<void> {
        const params = new URLSearchParams({
            name: element.name,
        });

        const path = element.type === 'animation' ? '/animations/play' : '/statics/set';

        await useFetch(`${apiUrl.value}${path}`, {
            method: 'POST',
            body: params,
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded'
            },
            mode: "no-cors"
        });
    }

    return {
        apply,
        getElements
    }
}