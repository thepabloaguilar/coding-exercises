package golang

func Merge(nums1 []int, m int, nums2 []int, n int) {
	// Usamos o `idx` para sabermos onde estamos inserindo no num1
	idx := m + n - 1
	m--
	n--

	// Essa solução foca em colocar todos os items de `nums2` dentro de `nums1`,
	// por isso a condição de parada é quando não há mais itens restantes de `nums2`
	// para colocar em `nums1`.
	for n >= 0 {
		// Verificamos se não ultrapassamos o limite do `nums1`, o tracking é feito pela variável `m`.
		// Se não ultrapassamos o limite e a posição de `nums1` for maior que a posição de `nums2` então
		// colocamos esse valor maior na posição `idx` em `nums1`.
		// Caso alguma dessas condições não seja verdadeira então podemos considerar que o valor desejado
		// para colocar em `nums[idx]` é o que está em `nums2[n]`.
		if m >= 0 && nums1[m] > nums2[n] {
			nums1[idx] = nums1[m]
			m--
		} else {
			nums1[idx] = nums2[n]
			n--
		}

		idx--
	}
}
