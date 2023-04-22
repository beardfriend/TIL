class Q:
    def __init__(self):
        self.queue = []
        
    def enqueue(self, val):
        self.queue.append(val)
    
    def dequeue(self):
        if len(self.queue) == 0:
            return None
        item = self.queue[0]
        self.queue = self.queue[1:]
        return item
    

q = Q()

q.enqueue(1)
q.enqueue(2)
q.enqueue(5)
print(q.dequeue())
print(q.dequeue())
print(q.dequeue())
print(q.dequeue())