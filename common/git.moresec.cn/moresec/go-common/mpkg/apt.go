package mpkg

/*
#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>
#define order(x) ((x) == '~' ? -1    \
        : isdigit((x)) ? 0   \
        : !(x) ? 0           \
        : isalpha((x)) ? (x) \
        : (x) + 256)
int CmpFragment(const char *A,const char *AEnd, const char *B,const char *BEnd)
{
    if (A >= AEnd && B >= BEnd)
        return 0;
    if (A >= AEnd)
    {
        if (*B == '~') return 1;
        return -1;
    }
    if (B >= BEnd)
    {
        if (*A == '~') return -1;
        return 1;
    }
    const char *lhs = A;
    const char *rhs = B;
    while (lhs != AEnd && rhs != BEnd)
    {
        int first_diff = 0;

        while (lhs != AEnd && rhs != BEnd &&
                (!isdigit(*lhs) || !isdigit(*rhs)))
        {
            int vc = order(*lhs);
            int rc = order(*rhs);
            if (vc != rc)
                return vc - rc;
            lhs++; rhs++;
        }

        while (*lhs == '0')
            lhs++;
        while (*rhs == '0')
            rhs++;
        while (isdigit(*lhs) && isdigit(*rhs))
        {
            if (!first_diff)
                first_diff = *lhs - *rhs;
            lhs++;
            rhs++;
        }

        if (isdigit(*lhs))
            return 1;
        if (isdigit(*rhs))
            return -1;
        if (first_diff)
            return first_diff;
    }

    // The strings must be equal
    if (lhs == AEnd && rhs == BEnd)
        return 0;

    // lhs is shorter
    if (lhs == AEnd)
    {
        if (*rhs == '~') return 1;
        return -1;
    }

    // rhs is shorter
    if (rhs == BEnd)
    {
        if (*lhs == '~') return -1;
        return 1;
    }

    // Shouldnt happen
    return 1;
}
int DoCmpVersion(const char *A,const char *AEnd, const char *B,const char *BEnd)
{
    // Strip off the epoch and compare it
    const char *lhs = A;
    const char *rhs = B;
    for (;lhs != AEnd && *lhs != ':'; lhs++);
    for (;rhs != BEnd && *rhs != ':'; rhs++);
    if (lhs == AEnd)
        lhs = A;
    if (rhs == BEnd)
        rhs = B;

    // Special case: a zero epoch is the same as no epoch,
    // so remove it.
    if (lhs != A)
    {
        for (; *A == '0'; ++A);
        if (A == lhs)
        {
            ++A;
            ++lhs;
        }
    }
    if (rhs != B)
    {
        for (; *B == '0'; ++B);
        if (B == rhs)
        {
            ++B;
            ++rhs;
        }
    }

    // Compare the epoch
    int Res = CmpFragment(A,lhs,B,rhs);
    if (Res != 0)
        return Res;

    // Skip the :
    if (lhs != A)
        lhs++;
    if (rhs != B)
        rhs++;

    // Find the last -
    const char *dlhs = AEnd-1;
    const char *drhs = BEnd-1;
    for (;dlhs > lhs && *dlhs != '-'; dlhs--);
    for (;drhs > rhs && *drhs != '-'; drhs--);

    if (dlhs == lhs)
        dlhs = AEnd;
    if (drhs == rhs)
        drhs = BEnd;

    // Compare the main version
    Res = CmpFragment(lhs,dlhs,rhs,drhs);
    if (Res != 0)
        return Res;

    // Skip the -
    if (dlhs != lhs)
        dlhs++;
    if (drhs != rhs)
        drhs++;

    return CmpFragment(dlhs,AEnd,drhs,BEnd);
}
*/
import "C"
import "unsafe"

// ubuntu/debian 包版本比对.
func CompareVersion(v1, v2 string) int {
	version1 := C.CString(v1)
	version2 := C.CString(v2)
	defer C.free(unsafe.Pointer(version1))
	defer C.free(unsafe.Pointer(version2))

	v1End := (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(version1)) + uintptr(len(v1))))
	v2End := (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(version2)) + uintptr(len(v2))))

	result := C.DoCmpVersion(version1, v1End, version2, v2End)
	return int(result)
}
