// Code generated by goderive DO NOT EDIT.

package set

func deriveSet(list []int) map[int]struct{} {
	set := make(map[int]struct{}, len(list))
	for _, v := range list {
		set[v] = struct{}{}
	}
	return set
}
