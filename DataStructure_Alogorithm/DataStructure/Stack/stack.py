class Stack:
    def __init__(self, max_size=5):
        self.stack = []
        self.top = -1
        self.max = max_size

    def isFull(self):
        if self.top == self.max:
            return True
        else:
            return False
    
    def isEmpty(self):
        if self.top == -1:
            return True
        else:
            return False
    
    def Push(self, val):
        if self.isFull():
            print("is Full")
            return
        else:
            self.stack.append(val)
            self.top += 1

    def Pop(self):
        if self.isEmpty():
            print("is empty")
            return
        else:
            item = self.stack[self.top]
            del self.stack[self.top]
            self.top -= 1
            return item
        

stack = Stack(10)

stack.Push(2)
stack.Push(3)
stack.Push(4)
print(stack.Pop())
print(stack.Pop())
print(stack.Pop())
print(stack.Pop())
            