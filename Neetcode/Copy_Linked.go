package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// Ваш вопрос - почему мы добавляем Val в map сразу, а не потом, но на самом деле в map мы сохраняем соответствие "оригинальный *Node" -> "новый *Node".
// Переделывать на "Val -> Node" нельзя, потому что Val не уникален и не отражает структуру списка и связей Random/Next.
// Вот чуть подробнее с пояснением:

func copyRandomList(head *Node) *Node {
	// Создаём отображение: оригинальный *Node -> новый *Node (deep copy)
	nodeMap := make(map[*Node]*Node)

	// Первый проход: копируем сами узлы (только Val)
	for cur := head; cur != nil; cur = cur.Next {
		nodeMap[cur] = &Node{Val: cur.Val}
	}

	// Второй проход: расставляем связи Next и Random у новых узлов
	for cur := head; cur != nil; cur = cur.Next {
		if nodeMap[cur] != nil {
			nodeMap[cur].Next = nodeMap[cur.Next]
			nodeMap[cur].Random = nodeMap[cur.Random]
		}
	}

	return nodeMap[head]
}
