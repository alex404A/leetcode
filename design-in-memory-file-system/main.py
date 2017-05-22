class File(object):

    def __init__(self, name, content):
        self.name = name
        self.content = content

    def addContent(self, content):
        self.content = self.content if self.content is not None else ''
        self.content += content

class Dir(object):
    def __init__(self, relativePath):
        self.relativePath = relativePath
        self.files = {}
        self.dirs = {}

    def getList(self, path):
        if len(path) == 1:
            fileObj = self.files.get(path[0])
            if fileObj is not None:
                return [fileObj.name]
        if len(path) == 0:
            files = [item.name for item in self.files.itervalues()]
            dirs = [item.relativePath for item in self.dirs.itervalues()]
            return sorted(dirs + files)
        relativePath = path[0]
        subDir = self.dirs.get(relativePath)
        return subDir.getList(path[1:])

    def makeDir(self, path):
        relativePath = path[0]
        subDir = self.dirs.get(relativePath)
        if len(path) == 1 and subDir is None:
            self.dirs[relativePath] = Dir(relativePath)
        elif len(path) > 1:
            if subDir is None:
                subDir = Dir(relativePath)
                self.dirs[relativePath] = subDir
            subDir.makeDir(path[1:])

    def addContent(self, path, content):
        if len(path) == 1:
            fileName = path[0]
            fileObj = self.files.get(fileName)
            if  fileObj is None:
                self.files[fileName] = File(fileName, content)
            else:
                fileObj.addContent(content)
        elif len(path) > 1:
            relativePath = path[0]
            subDir = self.dirs.get(relativePath)
            if subDir is None:
                subDir = Dir(relativePath)
                self.dirs[relativePath] = subDir
            subDir.addContent(path[1:], content)

    def readContent(self, path):
        if len(path) == 1:
            fileObj = self.files.get(path[0])
            return fileObj.content
        elif len(path) > 1:
            relativePath = path[0]
            subDir = self.dirs.get(relativePath)
            return subDir.readContent(path[1:])

class FileSystem(object):

    def __init__(self):
        self.root = Dir('.')

    def ls(self, path):
        """
        :type path: str
        :rtype: List[str]
        """
        return self.root.getList(self.getPathList(path))

    def mkdir(self, path):
        """
        :type path: str
        :rtype: void
        """
        self.root.makeDir(self.getPathList(path))

    def addContentToFile(self, filePath, content):
        """
        :type filePath: str
        :type content: str
        :rtype: void
        """
        self.root.addContent(self.getPathList(filePath), content)

    def readContentFromFile(self, filePath):
        """
        :type filePath: str
        :rtype: str
        """
        return self.root.readContent(self.getPathList(filePath))

    def getPathList(self, path):
        if path == '/':
            return []
        else:
            return path.split('/')[1:]

if __name__ == '__main__':
    obj = FileSystem()
    commands = ["mkdir","ls","ls","mkdir","ls","ls","addContentToFile","ls","ls","ls"]
    params = [["/goowmfn"],["/goowmfn"],["/"],["/z"],["/"],["/"],["/goowmfn/c","shetopcy"],["/z"],["/goowmfn/c"],["/goowmfn"]]
    def help (commands, params):
        length = len(commands)
        for i in range(len(commands)):
            result = getattr(obj, commands[i])(*params[i])
            if commands[i] == 'ls':
                print('files: ' + '/'.join(result))
            elif commands[i] == 'readContentFromFile':
                print('content: ' + result)
    help(commands, params)
